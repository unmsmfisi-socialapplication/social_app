import React, { useState } from 'react';

const Modal: React.FC = () => {
  const [isModalOpen, setModalOpen] = useState(false);

  const openModal = () => {
    setModalOpen(true);
  };

  const closeModal = () => {
    setModalOpen(false);
  };

  return (
    <div>
      <button onClick={openModal}>Open Modal</button>

      {isModalOpen && (
        <div className="modal">
          <div className="modal-content">
            <span className="close" onClick={closeModal}>
              &times;
            </span>
            <h2>Modal Content</h2>
            {/*STARTING VIEW*/}
          </div>
        </div>
      )}
    </div>
  );
};

export default Modal;
