import React, { useState } from 'react';
import '../styles/register-body_create_perfil.css';
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Poppins&display=swap" />

function ProfilePictureForm() {
  const [selectedImage, setSelectedImage] = useState(null);
  const [imageLoaded, setImageLoaded] = useState(false);

  const handleImageChange = (e) => {
    const file = e.target.files[0];
    setSelectedImage(URL.createObjectURL(file));
    setImageLoaded(true);
  };

  const handleSubmitAccepter = (e) => {
    e.preventDefault();
    // Lógica
  };

  const handleSubmitSauter = (e) => {
    e.preventDefault();
    // Lógica
  };

  return (
    <div className="container">
      <div className={`row ${imageLoaded ? 'image-loaded' : ''}`}>
        <div className={`column ${imageLoaded ? 'has-image' : ''}`}>
          <label htmlFor="imagen" className="upload-button">
            {selectedImage ? (
              <img src={selectedImage} alt="Foto de perfil" className="profile-image" />
            ) : (
              <span>Add a photo here +</span>
            )}
            <input
              type="file"
              id="imagen"
              name="imagen"
              accept="image/*"
              onChange={handleImageChange}
            />
          </label>
        </div>
        <div className={`column ${imageLoaded ? 'hidden' : ''}`}>
          <div className='text-container'>
            <h1>Configura una foto de perfil</h1>
            <p>Lorem Ipsum is simply dummy text of the printing and typesetting industry.</p>
          </div>
        </div>
      </div>
      <div className="button-container">
        {imageLoaded ? (
          <button type="submit" onClick={handleSubmitAccepter} className='btn-accepter'>Aceptar</button>
        ) : (
          <button type="submit" onClick={handleSubmitSauter} className='btn-sauter'>Deseo omitir este paso por ahora &gt;</button>
        )}
      </div>
    </div>
  );
}

export default ProfilePictureForm;
