import { StyleSheet } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { ThemeProvider } from "../ThemeProvider";

interface BaseLayoutProps {
  children: React.ReactNode;
}

const BaseLayout = ({ children }: BaseLayoutProps) => {
  return (
    <ThemeProvider>
      <SafeAreaView className="flex-1 bg-background" edges={['top', 'right', 'left']}>
        {children}
      </SafeAreaView>
    </ThemeProvider>
  );
};

export default BaseLayout;

const styles = StyleSheet.create({});
