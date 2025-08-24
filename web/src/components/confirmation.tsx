import { useNavigate, useParams } from "react-router-dom";
import { Button } from "./ui/button";
import axiosInstance from "@/api/axios-instance";
import { toast } from "sonner";

export const ConfirmationPage = () => {
 const {token = ''} = useParams();
 const redirect = useNavigate();

  const handleConfirm = async () => {
      try {
        console.log("Confirmation action triggered");
        const response = await axiosInstance.put("user/activate/" + token);
        console.log("Response:", response);
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
    <div>
      <h1>ConfirmationPage</h1>
      <Button onClick={() => handleConfirm()}>Click to confirm</Button>
    </div>
  );
};
