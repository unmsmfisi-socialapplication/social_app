import React, { Suspense } from "react";
export const ChatModule = React.lazy(() => import("./shared/components/chat"));

import "./App.css";
import Notification from "./shared/components/notification";

function App() {
  return (
    <>
      {/* Crear un dasboard */}
      {/* Usa Suspense para envolver el componente perezoso */}
      <Suspense fallback={<div>Cargando...</div>}>
        <ChatModule />
        <Notification
          avatarUrl="https://depor.com/resizer/kVf6CgNlGkGhV06Cj2mCGGUtb9w=/580x330/smart/filters:format(jpeg):quality(75)/cloudfront-us-east-1.images.arcpublishing.com/elcomercio/EFICHXCDBNDF3BVX2CILBGB6BQ.jpg"
          username="Sebastian"
          timeAgo="2h"
          action="le dio me gusta a tu publicación"
        />
      </Suspense>
    </>
  );
}

export default App;
