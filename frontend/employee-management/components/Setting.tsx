"use client";

import { Button } from "./ui/button";
import { useState } from "react";
import Profile from "@/components/Setting/Profile";
import Appearance from "@/components/Setting/App";
import Notifications from "@/components/Setting/Notifications";

export default function Setting() {
    const [menu, setMenu] = useState("Profile");
    return (
        <div className=" flex flex-col space-y-5">
            <div className=" flex flex-col space-y-2">
                <h1 className=" text-2xl font-bold">Settings</h1>
                <p className=" text-gray-700 dark:text-gray-300">
                    Manage your account and settings 
                </p>
            </div>
            <hr className=" border" />
            <div className=" flex flex-row max-w-screen">
                <div className=" w-[15vw] space-y-1">
                    <Button onClick={()=>{setMenu("Profile")}} variant="ghost" className={` w-full ${ menu === 'Profile' ? " dark:bg-blue-300/10 hover:bg-blue-100 hover:text-blue-500 text-blue-500 bg-blue-100" : ""}`}>Profile</Button>
                    <Button onClick={()=>{setMenu("Appearance")}} variant="ghost" className={` w-full ${ menu === 'Appearance' ? " dark:bg-blue-300/10 hover:bg-blue-100 hover:text-blue-500 text-blue-500 bg-blue-100" : ""}`}>Appearance</Button>
                    <Button onClick={()=>{setMenu("Notifications")}} variant="ghost" className={` w-full ${ menu === 'Notifications' ? " dark:bg-blue-300/10 hover:bg-blue-100 hover:text-blue-500 text-blue-500 bg-blue-100" : ""}`}>Notifications</Button>
                    

                </div>
                {
                    menu === "Profile" && (
                        <Profile />
                    )

                }
                {
                    menu === "Appearance" && (
                        <Appearance />
                    )
                }
                {
                    menu === "Notifications" && (
                        <Notifications />
                    )
                }

            </div>

        </div>
    )
}
