"use client";

import { Navigation } from "@/components/Shadcnui/Navigator";
import { MenuBar } from "./Shadcnui/MenuBar";
import { AvatarImpl } from "./Shadcnui/Avatar";
  
export default function Header() {
    return (
        <div className=" border border-b flex flex-row justify-between px-10 h-[7vh] items-center">
            <div className=" flex flex-row justify-center items-center space-x-7  ">
                <h1>
                    LOGO
                </h1>
                <Navigation />
            </div>
            <div className=" flex flex-row justify-center items-center space-x-7  ">
                <MenuBar />
                <AvatarImpl />
            </div>
        </div>

    )
}