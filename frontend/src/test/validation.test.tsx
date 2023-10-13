import { validateName, validateEmail, validatePassword } from '../utilities/Validation';

describe('validateName', () => {
  it('should return true for valid names', () => {
    expect(validateName('John Doe')).toBe(true);
    expect(validateName('María López')).toBe(true);
  });

  it('should return false for invalid names', () => {
    expect(validateName('123ABC')).toBe(false);
    expect(validateName('John@Doe')).toBe(false);
  });
  it('should return false for an empty string on name input', () => {
    expect(validateEmail('')).toBe(false);
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
  
  it('should return false for an empty string on email input', () => {
    expect(validateEmail('')).toBe(false);
  });
});

describe('validatePassword', () => {
  it('should return true for a valid password with at least 8 characters, 1 uppercase letter, 1 lowercase letter, 1 number, and 1 special character', () => {
    expect(validatePassword("Abcdefg1!")).toBe(true);
    expect(validatePassword("P@ssw0rd")).toBe(true);
    expect(validatePassword("Test123!")).toBe(true);
  });

  it('should return false for a password with less than 8 characters', () => {
    expect(validatePassword("Abcd1!")).toBe(false);
    expect(validatePassword("Abcdef!")).toBe(false);
  });

  it('should return false for a password without an uppercase letter', () => {
    expect(validatePassword("abcdefg1!")).toBe(false);
    expect(validatePassword("password1!")).toBe(false);
  });

  it('should return false for a password without a lowercase letter', () => {
    expect(validatePassword("ABCDEFG1!")).toBe(false);
    expect(validatePassword("PASSWORD1!")).toBe(false);
  });

  it('should return false for a password without a number', () => {
    expect(validatePassword("Abcdefg!")).toBe(false);
    expect(validatePassword("Password!")).toBe(false);
  });

  it('should return false for a password without a special character', () => {
    expect(validatePassword("Abcdefg1")).toBe(false);
    expect(validatePassword("Password1")).toBe(false);
  });

  it('should return false for an empty string', () => {
    expect(validatePassword("")).toBe(false);
  });
});

describe('validateUsername', () => {
  it('debería devolver true para correos válidos', () => {
    expect(validateEmail('wilfredohg57@gmail.com')).toBe(true);
    expect(validateEmail('devcell@gmail.com')).toBe(true);
  });

  it('debería devolver false para correos no válidos', () => {
    expect(validateEmail('12@3ABC')).toBe(false);
    expect(validateEmail('JohnWick')).toBe(false);
  });
  
  it('should return false for an empty string on email input', () => {
    expect(validateEmail('')).toBe(false);
  });
});