"use client"

import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select"
import { useState } from "react"
import { useTheme } from "next-themes"
import { useToast } from "@/components/ui/use-toast"

export default function Page(){
    const { toast } = useToast()
    const { setTheme } = useTheme()
    const [themes, setThemes] = useState("")
    const [font, setFont] = useState("")
    const onSubmit = () => {
        if (themes === "" || font === "") {
            toast({
                title: "Error",
                description: "Please select a theme and font",
              })
        }else{
            setTheme(themes)
            document.documentElement.style.setProperty("--font-sans", font)
            
            toast({
                title: "Success",
                description: "Preference updated successfully",
              })
        }
        
    }
    
    return(
        <div className=" w-[50vw] px-5 space-y-5">
            <div className=" flex flex-col">
                <h1 className=" text-lg font-medium">
                    Appearance
                </h1>
                <div className=" text-gray-400 dark:text-gray-400">
                    Customize the appearance of the app. Automatically switch between day and night themes.
                </div>
            </div>
            <hr />
            <div className="w-2/3 space-y-10">
                <div className="grid w-full max-w-sm items-center gap-1.5 space-y-2">
                    <Label htmlFor="font">Font</Label>
                    <Select
                        value={font}
                        onValueChange={setFont}                        
                    >
                    <SelectTrigger  className="w-[180px]">
                        <SelectValue placeholder="Select a font" />
                    </SelectTrigger>
                    <SelectContent >
                        <SelectGroup >
                        <SelectItem value="Inter">Inter</SelectItem>
                        <SelectItem value="Roboto">Roboto</SelectItem>
                        <SelectItem value="Kanit">Kanit</SelectItem>
                        </SelectGroup>
                    </SelectContent>
                    </Select>
                    <div className=" text-xs text-gray-500">
                        The font used in the app
                    </div>
                </div>
                <div className="grid w-full max-w-sm items-center gap-1.5 space-y-2">
                    <Label htmlFor="themes">Theme</Label>
                    <div className=" text-xs text-gray-500">
                        The theme used in the app
                    </div>
                    <div className=" flex flex-row space-x-4">
                        <div onClick={()=>{setThemes("light")}} className=" flex flex-col space-y-2 justify-center items-center">
                            <div className={` cursor-pointer w-[215px] ring-1 rounded-md ${themes === "light" ? "ring-gray-700 dark:ring-gray-100" : "ring-gray-200 dark:ring-gray-700"}  flex justify-center items-center h-[165px]`}>
                                <div className=" hover:bg-gray-200 p-3 flex flex-col justify-center items-center space-y-2  w-[200px] rounded-md bg-gray-200 h-[150px]">
                                    <div className=" w-[180px] space-y-2 bg-white h-[50px] rounded-md  flex flex-col justify-center items-start px-2">
                                        <div className=" w-[140px] h-[15px]  bg-gray-300 rounded-md">
                                        </div>
                                        <div className=" w-[160px] h-[15px]  bg-gray-300 rounded-md">
                                        </div>
                                    </div>
                                    <div className=" w-[180px] space-x-3 bg-white h-[30px] rounded-md flex flex-row items-center px-2">
                                        <div className=" w-[20px] h-[20px] rounded-full bg-gray-300">
                                        </div>
                                        <div className=" w-[120px] h-[15px]  bg-gray-300 rounded-md">
                                        </div>
                                    </div>
                                    <div className=" w-[180px] space-x-3 bg-white h-[30px] rounded-md flex flex-row items-center px-2">
                                        <div className=" w-[20px] h-[20px] rounded-full bg-gray-300">
                                        </div>
                                        <div className=" w-[120px] h-[15px]  bg-gray-300 rounded-md">
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <h1>
                                Light
                            </h1>
                        </div>
                        <div onClick={()=>{setThemes("dark")}} className=" flex flex-col space-y-2 justify-center items-center">
                            <div className={` cursor-pointer w-[215px] ring-1 rounded-md ${themes === "dark" ? "ring-gray-700 dark:ring-gray-100" : "ring-gray-200 dark:ring-gray-700"}  flex justify-center items-center h-[165px]`}>
                                <div className=" hover:bg-black p-3 flex flex-col justify-center items-center space-y-2  w-[200px] rounded-md bg-black h-[150px]">
                                    <div className=" w-[180px] space-y-2 bg-gray-700 h-[50px] rounded-md  flex flex-col justify-center items-start px-2">
                                        <div className=" w-[140px] h-[15px]  bg-gray-300 rounded-md">
                                        </div>
                                        <div className=" w-[160px] h-[15px]  bg-gray-300 rounded-md">
                                        </div>
                                    </div>
                                    <div className=" w-[180px] space-x-3 bg-gray-700 h-[30px] rounded-md flex flex-row items-center px-2">
                                        <div className=" w-[20px] h-[20px] rounded-full bg-gray-300">
                                        </div>
                                        <div className=" w-[120px] h-[15px]  bg-gray-300 rounded-md">
                                        </div>
                                    </div>
                                    <div className=" w-[180px] space-x-3 bg-gray-700 h-[30px] rounded-md flex flex-row items-center px-2">
                                        <div className=" w-[20px] h-[20px] rounded-full bg-gray-300">
                                        </div>
                                        <div className=" w-[120px] h-[15px]  bg-gray-300 rounded-md">
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <h1>
                                Dark
                            </h1>
                        </div>

                    </div>
                </div>  
                <Button onClick={onSubmit}>Update preference</Button>
            </div>
        </div>
    )
}