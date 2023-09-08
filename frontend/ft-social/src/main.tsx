import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { Switch } from '@mui/material'
import { Route } from 'react-router-dom'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <Switch>
      <Route path='/modal'>
        
      </Route>
    </Switch>
    <App />
  </React.StrictMode>,
)
