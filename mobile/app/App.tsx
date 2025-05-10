import { StatusBar } from "expo-status-bar";
import { Image, Pressable, Text, View } from "react-native";
import "../global.css";
import Hero2 from "../assets/hero2.png";
import BaseLayout from "../components/Layouts/BaseLayout";
import { useTheme } from "../components/ThemeProvider";

export const App = () => {
  const { toggleTheme } = useTheme();

  return (
    <BaseLayout>
            <StatusBar value="auto" />

      <View className="flex-1 bg-black justify-between items-center px-6 py-10">
        <View className="w-full flex-row items-center justify-start mt-4 mb-2">
          <Text className="text-[#7B61FF]/90 text-6xl font-extrabold font-['Poppins']">
            Lumina
          </Text>
        </View>

        <Image
          source={Hero2}
          className="w-92 h-92 mb-6 pr-0"
          resizeMode="contain"
        />

        <View className="w-full mb-8 gap-4">
          <Text className="w-80 h-50 justify-start text-white text-3xl font-extrabold font-['Poppins']">
            Jump start your solana business with Lumina
          </Text>
          <Text className="text-secondary-foreground font-medium text-lg mt-4">
            Empowering Indian Small Businesses with Seamless Solana-Powered
            Mobile Payment App
          </Text>
        </View>

        <Pressable
          className="w-full bg-primary py-4 rounded-2xl items-center mb-2 shadow-lg"
          style={{
            shadowColor: "#7B61FF",
            shadowOpacity: 0.8,
            shadowRadius: 10,
            shadowOffset: { width: 0, height: 4 },
          }}
        >
          <Text className="text-white text-lg font-semibold font-['Poppins']">
            Get Started
          </Text>
        </Pressable>
      </View>
    </BaseLayout>
  );
};
