"use client"

import * as React from "react"
import { addDays, format } from "date-fns"
import { Calendar as CalendarIcon } from "lucide-react"
import { DateRange } from "react-day-picker"

import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import { Calendar } from "@/components/ui/calendar"
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover"
import dayjs from "dayjs"

interface DatePickerRangeProps extends React.HTMLAttributes<HTMLDivElement> {
  onStartDateChange: Function,
  onEndDateChange: Function
}

export default function DatePickerRange({
  className, onStartDateChange, onEndDateChange
}: DatePickerRangeProps) {
  const [date, setDate] = React.useState<DateRange | undefined>({
    from: new Date(),
    to: addDays(new Date(), 3),
  })

  const handleDateChange = (range: DateRange | undefined) => {
    setDate(range);
    onStartDateChange(range?.from);
    onEndDateChange(range?.to);
  }

  React.useEffect(() => {
    onStartDateChange(date?.from);
    onEndDateChange(date?.to);
  }, [])

  return (
    <div className={cn("grid gap-2", className)}>
      <Popover>
        <PopoverTrigger asChild>
          <Button
            id="date"
            variant={"outline"}
            className={cn(
              "w-full justify-center text-left font-normal",
              !date && "text-muted-foreground"
            )}
          >
            <CalendarIcon className="mr-2 h-4 w-4" />
            {date?.from ? (
              date.to ? (
                <>
                  {format(date.from, "LLL dd, y")} -{" "}
                  {format(date.to, "LLL dd, y")}
                </>
              ) : (
                format(date.from, "LLL dd, y")
              )
            ) : (
              <span>Pick a date</span>
            )}
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-auto p-0" align="start">
          <Calendar
            initialFocus
            mode="range"
            defaultMonth={date?.from}
            selected={date}
            onSelect={handleDateChange}
            numberOfMonths={2}
          />
        </PopoverContent>
      </Popover>
    </div>
  )
}
