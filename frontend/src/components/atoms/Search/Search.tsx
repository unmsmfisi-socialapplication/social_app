import React from 'react'
import { TextField, InputAdornment, IconButton } from '@mui/material'
import SearchIcon from '@mui/icons-material/Search'
import './Search.scss' // Importa el archivo de estilos SCSS

const WSearch = () => {
    return (
        <TextField
            variant="standard"
            placeholder="Buscar" // Agrega un marcador de posición
            className="custom-search" // Agrega una clase personalizada
            InputProps={{
                startAdornment: (
                    <InputAdornment position="start">
                        <IconButton>
                            <SearchIcon />
                        </IconButton>
                    </InputAdornment>
                ),
            }}
        />
    )
}

export default WSearch
