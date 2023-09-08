import logo from './logo.svg';
import './App.css';
import React, { useState } from 'react';


function App() {
  const [nombre, setNombre] = useState('');
  
  const handleNombreChange = (e) => {
    setNombre(e.target.value);
  };
  return (
    <div className="App">
      <div className="Izquierdo">
          <div className="perfil-container">
              <h2 className="titulo">Crea tu perfil</h2>
              <div className="perfil-circle"></div>
          </div>
          <div className="formulario-container">
            <h2 className="NamePerfil">Ahora, ingresa tu nombre de perfil</h2>
              <form>
                <div className="form-group">
                  <input
                      className="campo"
                      placeholder='Nombre'
                      type="text"
                      id="nombre"
                      value={nombre}
                      onChange={handleNombreChange}
                      required
                  />
                </div>
              </form>
          </div>
      </div>
      <div className="boton-container">
        <button className="boton-hecho">Hecho</button>
      </div>
    </div>
  );
}

export default App;
