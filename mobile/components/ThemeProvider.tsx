import React, { createContext, useContext } from "react";
import { View } from "react-native";
import { useColorScheme } from "nativewind";

interface ThemeProviderProps {
  children: React.ReactNode;
}

interface ThemeContextType {
  theme: "light" | "dark";
  toggleTheme: () => void;
}

export const ThemeContext = createContext<ThemeContextType>({
  theme: "dark",
  toggleTheme: () => {},
});

export const useTheme = () => useContext(ThemeContext);

export const ThemeProvider = ({ children }: ThemeProviderProps) => {
  const { colorScheme = "dark", setColorScheme } = useColorScheme();

  const toggleTheme = () => {
    setColorScheme(colorScheme === "dark" ? "light" : "dark");
  };

  return (
    <ThemeContext.Provider value={{ theme: colorScheme, toggleTheme }}>
      <View className={`flex-1 ${colorScheme}`}>
        {children}
      </View>
    </ThemeContext.Provider>
  );
}; 