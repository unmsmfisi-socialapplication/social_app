"use client";
import { useState } from "react";
import Layout from "../layout";
import { Button } from "@mui/material";
import { WButton } from "@/components";
import { WCircleIcon } from "@/components";
import CheckIcon from "@mui/icons-material/Check";

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
      <WCircleIcon iconSize={30} icon={CheckIcon} />
    </Layout>
  );
}
