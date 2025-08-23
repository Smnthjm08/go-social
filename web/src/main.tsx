import { createRoot } from "react-dom/client";
import "./index.css";
import { RouterProvider } from "react-router-dom";
import { ThemeProvider } from "./components/theme-provider.tsx";
import { Toaster } from "@/components/ui/sonner";
import { router } from "./router.tsx";

createRoot(document.getElementById("root")!).render(
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <RouterProvider router={router} />
      <Toaster />
      {/* <App /> */}
    </ThemeProvider>
);
