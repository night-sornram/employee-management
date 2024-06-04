import { Button } from "@/components/ui/button";
import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";
import customParseFormat from "dayjs/plugin/customParseFormat";
export default function DownloadDataPage () {
    return (
        <main className="py-[5%] px-[5%]  md:w-[80%] 2xl:w-[60%] flex flex-col gap-10">
            <div>
                <h1 className="font-bold text-2xl">
                    Download Data
                </h1>
            </div>
            <div className="self-center">
                <Button>
                    Download
                </Button>
            </div>
        </main>
    )
}