// import { API_URL } from "@/App";
import axios from "axios";
import { useNavigate, useParams } from "react-router-dom";
import { Button } from "./ui/button";

const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8000/api";

export const ConfirmationPage = () => {
 const {token = ''} = useParams();
 const redirect = useNavigate();

  const handleConfirm = async () => {
      try {
        console.log("Confirmation action triggered");
        const response = await axios.put(`${API_URL}/user/activate/${token}`);
        console.log("Response:", response);
        if (response.status === 200){
            // redirect to the "/" page or show a success message
            redirect("/");
            alert("Account activated successfully!");
        }
    } catch (error) {
        console.error("Error activating account:", error);
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
