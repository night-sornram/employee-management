'use client'
 
import { usePathname  , useRouter } from 'next/navigation'
import { Button } from "./ui/button";


export default function SettingSideBar({
    children
} : { children: React.ReactNode }) {
    const params = usePathname()
    const router = useRouter()
    return (
        <main className=" w-auto p-5 flex flex-col space-y-5">
            <div className=" flex flex-col space-y-2">
                <h1 className=" text-2xl font-bold">Settings</h1>
                <p className=" text-gray-700 dark:text-gray-300">
                    Manage your account and settings 
                </p>
            </div>
            <hr className=" border" />
            <div className=" flex flex-row" >
                <div className=" w-[15vw] space-y-1 flex flex-col">
                    <Button className={` ${ params === "/setting/profile" ? " dark:bg-blue-300/10 hover:bg-blue-100 hover:text-blue-500 text-blue-500 bg-blue-100  " : ""}   `} onClick={()=>{router.push("/setting/profile")}} variant="ghost">
                        Profile
                    </Button>
                    <Button className={` ${ params === "/setting/appearance" ? " dark:bg-blue-300/10 hover:bg-blue-100 hover:text-blue-500 text-blue-500 bg-blue-100  " : ""}   `} onClick={()=>{router.push("/setting/appearance")}} variant="ghost">
                        Appearance
                    </Button>
                    <Button className={` ${ params === "/setting/notifications" ? " dark:bg-blue-300/10 hover:bg-blue-100 hover:text-blue-500 text-blue-500 bg-blue-100  " : ""}   `} onClick={()=>{router.push("/setting/notifications")}} variant="ghost">
                        Notifications
                    </Button>
                    

                </div>
                {children}
                

            </div>

        </main>
    )
}
