
import { EnterIcon , ExitIcon , CounterClockwiseClockIcon , CalendarIcon  } from "@radix-ui/react-icons"

import { Button } from "@/components/ui/button"


export default function SideBar({children}: {children: React.ReactNode}) {
    return (
        <div className=" h-[93vh] w-screen flex flex-row">
            <div className=" w-[15vw] h-full border-r flex flex- p-5">
                <div className=" flex flex-col space-y-5 w-full">
                    <h1 className="   text-sm  ">
                        CHECK-IN CHECK-OUT
                    </h1>
                    <div className=" flex flex-col space-y-3 w-full">
                        <Button variant="ghost">
                            <EnterIcon className="mr-2 h-5 w-5" />
                            <h1 className="   ">
                                Check In
                            </h1> 
                        </Button>
                        <Button variant="ghost">
                            <ExitIcon  className="mr-2 h-5 w-5  " />
                            <h1 className="    ">
                                Check Out
                            </h1> 
                        </Button>
                        <Button variant="ghost">
                            <CounterClockwiseClockIcon  className="mr-2 h-5 w-5     " />
                            <h1 className="    ">
                                History
                            </h1> 
                        </Button>
                    </div>
                    <h1 className="   text-sm  ">
                        LEAVE MANAGEMENT
                    </h1>
                    <div className=" flex flex-col space-y-3 w-full">
                        <Button variant="ghost">
                            <CalendarIcon className="mr-2 h-5 w-5     " />
                            <h1 className="    ">
                                Request
                            </h1> 
                        </Button>
                        <Button variant="ghost">
                            <CounterClockwiseClockIcon  className="mr-2 h-5 w-5     " />
                            <h1 className="    ">
                                History
                            </h1> 
                        </Button>
                    </div>
                    
                    
                </div>
                

            </div>
            {children}
        </div>
    
    
    )}