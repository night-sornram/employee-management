import NavBar from "@/components/NavBar/NavBar";
import TopBar from "@/components/TopBar/TopBar";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
        <TopBar/>
        <div className="flex w-full">
            <NavBar/>
            {children}
        </div>
        
    </>
  );
}
