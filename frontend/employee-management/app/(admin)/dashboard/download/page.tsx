'use client'
import { Button } from "@/components/ui/button";
import DownloadAttendance from "@/lib/DownloadAttendance";
import { useSession } from "next-auth/react";
import { useEffect, useState } from "react";

export default function DownloadDataPage() {
    const { data: session } = useSession();
    const [downloadUrl, setDownloadUrl] = useState<string | null>();

    const handleDownload = async () => {
        if (!session) return;
        
        try {
            const blob = await DownloadAttendance(session.user.token);
            const url = URL.createObjectURL(blob);
            setDownloadUrl(url);
        } catch (error) {
            console.error("Error downloading attendance data:", error);
        }
    };

    useEffect(() => {
        // Cleanup the object URL when the component unmounts or when downloadUrl changes
        return () => {
            if (downloadUrl) {
                URL.revokeObjectURL(downloadUrl);
            }
        };
    }, [downloadUrl]);

    return (
        <main className="py-[5%] px-[5%] md:w-[80%] 2xl:w-[60%] flex flex-col gap-10">
            <div>
                <h1 className="font-bold text-2xl">Download Data</h1>
            </div>
            <div className="self-center">
                <Button onClick={handleDownload}>
                    Download
                </Button>
                {downloadUrl && (
                    <a href={downloadUrl} download="attendances.csv" style={{ display: 'none' }} ref={(el) => el?.click()}></a>
                )}
            </div>
        </main>
    );
}