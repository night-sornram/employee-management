import { Button } from "@/components/ui/button"
import { BellIcon  } from "@radix-ui/react-icons"

import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover"

export function MenuBar() {
  return (
    <Popover>
      <PopoverTrigger asChild>
        <button>
          <BellIcon className="mr-2 h-6 w-6 " />
        </button>
      </PopoverTrigger>
      <PopoverContent className="w-60 mr-20">
        <div className="grid gap-4">
          <div className="space-y-2">
            <h4 className="font-medium leading-none">Notifications</h4>
            <p className="text-sm text-muted-foreground">
              No Notifications
            </p>
          </div>
        </div>
      </PopoverContent>
    </Popover>
  )
}
