// Handles xsd:duration to time.Duration conversion

package xsd

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

// Extending time constants with the values for day, week, month, year
const (
	Day = time.Hour * 24
	// These values are not precise, probably why the time package does not implement them
	// We need them here because xsd:duration uses them
	Monthish = Day * 30
	Yearish  = Day * 356
)

// durationPair holds information about a pair of (uint/ufloat)(PeriodTag) values in a xsd:duration string
type durationPair struct {
	v   time.Duration
	typ byte
}

func getTimeBaseDuration(b byte) time.Duration {
	switch b {
	case tagHour:
		return time.Hour
	case tagMinute:
		return time.Minute
	case tagSecond:
		return time.Second
	}
	return 0
}

func getDateBaseDuration(b byte) time.Duration {
	switch b {
	case tagYear:
		return Yearish
	case tagMonth:
		return Monthish
	case tagDay:
		return Day
	}
	return 0
}

// validByteForFloats checks
func validByteForFloats(b byte) bool {
	//        +    ,    -    .     0 1 2 3 4 5 6 7 8 9
	return (b >= 43 && b <= 46) || (b >= 48 && b <= 57)
}

func parseTagWithValue(data []byte, start, tagPos int, isTime bool) (*durationPair, error) {
	d := new(durationPair)
	d.typ = data[tagPos]
	if d.typ == tagSecond {
		// seconds can be represented in float, we need to parse accordingly
		if v, err := strconv.ParseFloat(string(data[start:tagPos]), 32); err == nil {
			d.v = time.Duration(float64(time.Second) * v)
		}
	} else {
		if v, err := strconv.ParseInt(string(data[start:tagPos]), 10, 32); err == nil {
			if isTime {
				d.v = getTimeBaseDuration(d.typ) * time.Duration(v)
			} else {
				d.v = getDateBaseDuration(d.typ) * time.Duration(v)
			}
		}
	}
	if d.v < 0 {
		return nil, fmt.Errorf("the minus sign must appear first in the duration")
	}
	return d, nil
}

// loadUintVal receives the data and position at which to try to load the duration element
// isTime is used for distinguishing between P1M - 1 month, and PT1M - 1 minute
// it returns a duration pair if it can read one at the start of data,
//    and the number of bytes read from data
func loadUintVal(data []byte, start int, isTime bool) (*durationPair, int, error) {
	for i := start; i < len(data); i++ {
		chr := data[i]
		if validTag(data[i]) {
			d, err := parseTagWithValue(data, start, i, isTime)
			return d, i, err
		}
		if validByteForFloats(chr) {
			continue
		}
		return nil, i, fmt.Errorf("invalid character %c at pos %d", chr, i)
	}
	return nil, len(data), fmt.Errorf("unable to recognize any duration value")
}

const (
	tagDuration = 'P'
	tagTime     = 'T'
	tagYear     = 'Y'
	tagMonth    = 'M'
	tagDay      = 'D'
	tagHour     = 'H'
	tagMinute   = 'M'
	tagSecond   = 'S'
)

func validTag(b byte) bool {
	return b == tagYear || b == tagMonth || b == tagDay || b == tagHour || b == tagMinute || b == tagSecond
}

// Unmarshal takes a byte array and unmarshals it to a time.Duration value
// It is used to parse values in the following format: -PuYuMuDTuHuMufS, where:
//   * - shows if the duration is negative
//   * P is the duration tag
//   * T is the time tag separator
//   * Y,M,D,H,M,S are tags for year, month, day, hour, minute, second values
//   * u is an unsigned integer value
//   * uf is an unsigned float value (just for seconds)
func Unmarshal(data []byte, d *time.Duration) error {
	if len(data) == 0 {
		return fmt.Errorf("invalid xsd:duration: empty value")
	}
	pos := 0
	// loading if the value is negative
	negative := data[pos] == '-'
	if negative {
		// skipping over the minus
		pos++
	}
	if data[pos] != tagDuration {
		return fmt.Errorf("invalid xsd:duration: first character must be %q", 'P')
	}
	// skipping over the "P"
	pos++

	onePastEnd := len(data)
	if pos >= onePastEnd {
		return fmt.Errorf("invalid xsd:duration: at least one number and designator are required")
	}
	isTime := false
	duration := time.Duration(0)
	for {
		if data[pos] == tagTime {
			pos++
			isTime = true
		}
		p, cnt, err := loadUintVal(data, pos, isTime)
		if err != nil {
			return fmt.Errorf("invalid xsd:duration: %w", err)
		}
		duration += p.v
		pos = cnt + 1
		if pos+1 >= onePastEnd {
			break
		}
	}
	if negative {
		duration *= -1
	}
	if onePastEnd > pos+1 {
		return fmt.Errorf("data contains more bytes than we are able to parse")
	}
	if d == nil {
		return fmt.Errorf("unable to store time.Duration to nil pointer")
	}
	*d = duration
	return nil
}

func Days(d time.Duration) float64 {
	dd := d / Day
	h := d % Day
	return float64(dd) + float64(h)/(24*60*60*1e9)
}

func Months(d time.Duration) float64 {
	m := d / Monthish
	w := d % Monthish
	return float64(m) + float64(w)/(4*7*24*60*60*1e9)
}

func Years(d time.Duration) float64 {
	y := d / Yearish
	m := d % Yearish
	return float64(y) + float64(m)/(12*4*7*24*60*60*1e9)
}

func Marshal(d time.Duration) ([]byte, error) {
	if d == 0 {
		return []byte{tagDuration, tagTime, '0', tagSecond}, nil
	}

	neg := d < 0
	if neg {
		d = -d
	}
	y := Years(d)
	d -= time.Duration(y) * Yearish
	m := Months(d)
	d -= time.Duration(m) * Monthish
	dd := Days(d)
	d -= time.Duration(dd) * Day
	H := d.Hours()
	d -= time.Duration(H) * time.Hour
	M := d.Minutes()
	d -= time.Duration(M) * time.Minute
	s := d.Seconds()
	d -= time.Duration(s) * time.Second
	b := bytes.Buffer{}
	if neg {
		b.Write([]byte{'-'})
	}
	b.Write([]byte{'P'})
	if int(y) > 0 {
		b.WriteString(fmt.Sprintf("%d%c", int64(y), tagYear))
	}
	if int(m) > 0 {
		b.WriteString(fmt.Sprintf("%d%c", int64(m), tagMonth))
	}
	if int(dd) > 0 {
		b.WriteString(fmt.Sprintf("%d%c", int64(dd), tagDay))
	}

	if H+M+s > 0 {
		b.Write([]byte{tagTime})
		if int(H) > 0 {
			b.WriteString(fmt.Sprintf("%d%c", int64(H), tagHour))
		}
		if int(M) > 0 {
			b.WriteString(fmt.Sprintf("%d%c", int64(M), tagMinute))
		}
		if int(s) > 0 {
			if s-float64(int(s)) < 0.01 {
				b.WriteString(fmt.Sprintf("%d%c", int(s), tagSecond))
			} else {
				b.WriteString(fmt.Sprintf("%.1f%c", s, tagSecond))
			}
		}
	}
	return b.Bytes(), nil
}
