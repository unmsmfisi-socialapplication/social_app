'use client'
import React, { useState } from 'react';
import '../app/reactionPost/postStyles.scss';

const ReactionsComponent: React.FC = () => {
  const [likes, setLikes] = useState(0);
  const [hearts, setHearts] = useState(0);
  const [laughs, setLaughs] = useState(0);

  const handleLike = () => {
    setLikes(likes + 1);
  }

  const handleHeart = () => {
    console.log(hearts + 1);
    setHearts(hearts + 1);
  }

  const handleLaugh = () => {
    setLaughs(laughs + 1);
  }
  return (
    <div className="reactions">
      <button onClick={handleLike}>👍 {likes}</button>
      <button onClick={handleHeart}>❤️ {hearts}</button>
      <button onClick={handleLaugh}>😂 {laughs}</button>
    </div>
  );
};

export default ReactionsComponent;
