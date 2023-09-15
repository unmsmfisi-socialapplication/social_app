import React, { Suspense } from "react";
export const ChatModule = React.lazy(() => import("./shared/components/chat"));
import LoginPageContainer from "./containers/LoginPageContainer";

import "./App.css";

function App() {
  return (
    <>
      {/* TODO: Page routing */}
      <LoginPageContainer />
      {/* Crear un dashboard */}
      {/* Usa Suspense para envolver el componente perezoso */}
      <Suspense fallback={<div>Cargando...</div>}>
        <ChatModule />
      </Suspense>
    </>
  );
}

export default App;
