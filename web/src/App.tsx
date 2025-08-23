import { Button } from "@/components/ui/button";
import axios from "axios";
import { toast } from "sonner";

// export const API_URL = import.meta.VITE_API_URL || "http://localhost:8000/v1";

function App() {
  async function fetchHealth() {
    const res = await axios.get("http://localhost:8000/v1/health");
    toast.success(`status: ${res?.data?.status} | env: ${res?.data?.env}`);
    console.log(res.data);
  }

  return (
    <div className="flex min-h-svh flex-col items-center justify-center">
      <Button onClick={() => fetchHealth()}>Click me</Button>
    </div>
  );
}

export default App;
