import React, { Suspense } from 'react'
export const ChatModule = React.lazy(() => import('./shared/components/chat'));
import { Button } from './shared/components';

import './App.css'


function App() {


  return (
    <>
      {/* Crear un dasboard */}
      {/* Usa Suspense para envolver el componente perezoso */}
      <Suspense fallback={<div>Cargando...</div>}>
        <ChatModule />
        <Button />
      </Suspense>
    </>
  )
}

export default App
