import { validateName, validateEmail } from '../utilities/Validation';

describe('validateName', () => {
  it('debería devolver true para nombres válidos', () => {
    expect(validateName('John Doe')).toBe(true);
    expect(validateName('María López')).toBe(true);
  });

  it('debería devolver false para nombres no válidos', () => {
    expect(validateName('123ABC')).toBe(false);
    expect(validateName('John@Doe')).toBe(false);
  });
});

describe('validateEmail', () => {
  it('debería devolver true para correos válidos', () => {
    expect(validateEmail('wilfredohg57@gmail.com')).toBe(true);
    expect(validateEmail('devcell@gmail.com')).toBe(true);
  });

  it('debería devolver false para correos no válidos', () => {
    expect(validateEmail('12@3ABC')).toBe(false);
    expect(validateEmail('JohnWick')).toBe(false);
  });
});