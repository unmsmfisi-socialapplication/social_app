import { validarContraseñas } from '../utilities/Functions';
import { passwordCompare } from '../utilities/Constant';

describe('Validar Contraseñas', () => {
  it('Debería devolver nulo si las contraseñas coinciden', () => {
    const contraseña1 = 'contraseña123';
    const contraseña2 = 'contraseña123';
    const resultado = validarContraseñas(contraseña1, contraseña2);
    expect(resultado).toBeNull();
  });

  it('Debería devolver un mensaje de error si las contraseñas no coinciden', () => {
    const contraseña1 = 'contraseña123';
    const contraseña2 = 'otracontraseña';
    const resultado = validarContraseñas(contraseña1, contraseña2);
    expect(resultado).toBe(passwordCompare);
  });
});