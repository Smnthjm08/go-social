import { useNavigate, useParams } from "react-router-dom";
import { Button } from "../../components/ui/button";
import axiosInstance from "@/api/axios-instance";
import { toast } from "sonner";

export const ConfirmationPage = () => {
 const {token = ''} = useParams();
 const redirect = useNavigate();

  const handleConfirm = async () => {
      try {
        const response = await axiosInstance.put("auth/activate/" + token);
        if (response.status === 204){
            redirect("/");
            toast.success("Account activated successfully!");
        }
    } catch (error) {
        console.error("Error activating account:", error);
        toast.error("Failed to activate account. Please try again.");
        alert("Failed to activate account. Please try again.");
    }
  };

  return (
    <div className="flex min-h-svh w-full flex-col items-center justify-center p-6 md:p-10 gap-4">
      <h1>ConfirmationPage</h1>
      <Button onClick={() => handleConfirm()}>Click to confirm</Button>
    </div>
  );
};
