'use client'
import { useState } from "react"
import Layout from "../layout"
import { Button } from "@mui/material";
export default function TestPage() {
    const [count, setCount] = useState(0);

    const handleCount = () => {
        setCount(count + 1);
        alert(count);
    }
    return (
        <Layout>
            <Button variant="contained">Hello World</Button>
            <h1>Test Page</h1>
            <button onClick={handleCount}>presioname</button>
        </Layout>
      );
}