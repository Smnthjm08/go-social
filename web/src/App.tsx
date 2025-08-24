import { Button } from "@/components/ui/button";
import { toast } from "sonner";
import axiosInstance from "./api/axios-instance";
import { Navbar } from "./components/ui/io/navbar";

const App = () => {
  async function fetchHealth() {
    try {
      const res = await axiosInstance.get("http://localhost:8000/v1/health");
      toast.success(`status: ${res?.data?.status} | env: ${res?.data?.env}`);
      console.log(res.data);
    } catch (err) {
      toast.error("Failed to fetch health check");
      console.error(err);
    }
  }

  return (
    <div className="absolute inset-0 bg-[radial-gradient(125%_125%_at_50%_90%,white_40%,#6366f1_100%)]">
        <Navbar />
      <div className="min-h-screen w-full relative flex flex-col items-center justify-center">
        <h2 className="mb-4 text-3xl font-bold text-gray-800">
          Welcome to Go Social!
        </h2>
        <h4 className="text-gray-600 text-center pb-2">
          This is currently under development
        </h4>
        <Button onClick={fetchHealth}>Click me</Button>
      </div>
    </div>
  );
};

export default App;
