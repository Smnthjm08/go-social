import { Button } from "@/components/ui/button";
import axios from "axios";
import { useEffect } from "react";

// export const API_URL = import.meta.VITE_API_URL || "http://localhost:8000/v1";

function App() {

useEffect(() => {
  async function fetchHealth() {
    const res = await axios.get("http://localhost:8000/v1/health");
    console.log(res.data);
  }
  fetchHealth();
}, []);

  return (
      <div className="flex min-h-svh flex-col items-center justify-center">
        <Button>Click me</Button>
      </div>
  );
}

export default App;
