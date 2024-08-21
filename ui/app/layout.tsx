import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { NextUIProvider } from "@nextui-org/system";
import ReduxProvider from "@/redux/provider";
import Toast from "@/components/Toaster";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
    title: "Let's Live",
    description: "Powered by NextJS",
};

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang="en">
            <body className={inter.className}>
                <ReduxProvider>
                    <NextUIProvider>{children}</NextUIProvider>
                    <Toast />
                </ReduxProvider>
            </body>
        </html>
    );
}
