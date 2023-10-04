"use client";
import PersonIcon from '@mui/icons-material/Person';
import HideImageIcon from '@mui/icons-material/HideImage';
import React from 'react'
import "./index.scss"

//componente creado para ser llamado dentro de un formulario superior

interface AvatarInputProps {
  avatarDefaultURL?: string; 
  onChangeAvatar(avatar:File | undefined):void; // funcion para actualizar el valor del objeto que contendra los valores del formulario principal
  avatarValue:File | undefined;  // valor del archivo actual
}

const AvatarInput:React.FC<AvatarInputProps> = ({avatarDefaultURL,onChangeAvatar, avatarValue}) => {

  const handleChangeImage = (e:React.FormEvent<HTMLInputElement>) => {
    const target = e.target as HTMLInputElement & {
        files:FileList
    }
    if (target.files[0]) {
        onChangeAvatar(target.files[0]) //actualizamos el valor del objeto del formulario principal 
    } 
  }

  return (
    <div className='avatar_input_main_container'>
        <div className='avatars_container'>
            <div className='avatar_container' >
               {avatarDefaultURL? <img  src={avatarDefaultURL} />:<PersonIcon />}
            </div>
            <div className='avatar_container' >
                {avatarValue? <img src={URL.createObjectURL(avatarValue)}/>:<HideImageIcon/>}
            </div>
        </div>
        <label className='avatar_label_input'  htmlFor='avatar' >Editar foto</label>
        <input  className='avatar_input' id="avatar" name='avatar' type="file" accept="image/*" multiple={false}  onChange={handleChangeImage} />
    </div>
  )
}

export default AvatarInput

AvatarInput.defaultProps={
    avatarDefaultURL:"",
}