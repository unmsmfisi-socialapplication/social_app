'use client'
import { useState } from "react"
import Layout from "../layout"
import { Button } from "@mui/material";
import {TextField} from '@mui/material';
import { WButton, WInput } from "@/components";
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome'
import {faUser} from '@fortawesome/free-solid-svg-icons'
export default function TestPage() {
    const [count, setCount] = useState(0);

  const handleCount = () => {
    setCount(count + 1);
    alert(count);
  };
  return (
    <Layout>
      <Button variant="contained">Hello World</Button>
      <div
        style={{
          paddingLeft: "15px",
          width: "1000px",
          height: "100px",
          backgroundColor: "red",
        }}
      >
        <WButton>DD</WButton>
        <WButton typeColor="secondary">DD</WButton>
        <WButton typeColor="secondary" size="large">
          DD
        </WButton>
      </div>
      <h1>Test Page</h1>
      <button onClick={handleCount}>presioname</button>


      <div>
        TextField
      <WInput placeholder="Correo Electrónico" size="medium" icon={faUser} />
      </div>
    </Layout>

      
  );
}