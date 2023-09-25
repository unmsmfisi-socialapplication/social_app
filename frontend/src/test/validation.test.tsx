import { validateName } from '../utilities/Validation';

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