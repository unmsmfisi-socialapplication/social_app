import { validatePassword } from '../utilities/Functions';

describe('Validate passwords', () => {
  it('It should return null if the passwords match', () => {
    const password1 = 'password123';
    const password2 = 'password123';
    const result = validatePassword(password1, password2);
    expect(result).toBe(true);
  });

  it('It should return an error message if the passwords do not match', () => {
    const password1 = 'password123';
    const password2 = 'anotherpassword';
    const result = validatePassword(password1, password2);
    expect(result).toBe(false);
  });
});