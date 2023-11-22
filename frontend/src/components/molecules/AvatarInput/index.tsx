'use client'
import PersonIcon from '@mui/icons-material/Person'
import HideImageIcon from '@mui/icons-material/HideImage'
import React, { useState } from 'react'
import './index.scss'

//componente creado para ser llamado dentro de un formulario superior

interface AvatarInputProps {
    avatarDefaultURL?: string
    testId?: string
}

const AvatarInput: React.FC<AvatarInputProps> = ({ avatarDefaultURL, testId }) => {
    const [selectedImage, setSelectedImage] = useState<File | null>(null)
    const handleChangeImage = (e: React.FormEvent<HTMLInputElement>) => {
        const target = e.target as HTMLInputElement & {
            files: FileList
        }
        if (target.files[0]) {
            setSelectedImage(target.files[0]) //actualizamos el valor del objeto del formulario principal
        }
    }

    return (
        <div className="avatar_input_main_container">
            <div className="avatars_container">
                <div className="avatar_container">
                    {avatarDefaultURL ? <img src={avatarDefaultURL} alt="Default Image" /> : <PersonIcon />}
                </div>
                <div className="avatar_container">
                    {selectedImage ? (
                        <img src={URL.createObjectURL(selectedImage)} alt="Selected Image" />
                    ) : (
                        <HideImageIcon />
                    )}
                </div>
            </div>
            <label className="avatar_label_input" htmlFor="avatar">
                Editar foto
            </label>
            <input
                className="avatar_input"
                id="avatar"
                data-testid={testId}
                name="avatar"
                type="file"
                accept="image/*"
                multiple={false}
                onChange={handleChangeImage}
            />
        </div>
    )
}

export default AvatarInput

AvatarInput.defaultProps = {
    avatarDefaultURL: '',
}
