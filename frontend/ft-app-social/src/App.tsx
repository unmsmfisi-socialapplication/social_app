import React, { Suspense } from 'react'
export const ChatModule = React.lazy(() => import('./shared/components/chat'));

import './App.css'


function App() {


  return (
    <>
      {/* Crear un dasboard */}
      {/* Usa Suspense para envolver el componente perezoso */}
      <Suspense fallback={<div>Cargando...</div>}>
        <ChatModule />
      </Suspense>
    </>
  )
}

export default App
