"use client";

import { RootState, useAppSelector } from '@/store/store'


export default function FontProvider({ children} : {children: React.ReactNode , } ) : React.ReactNode {
    const data = useAppSelector((state: RootState) => state.fontSlice)
    document.body.style.fontFamily = `${data.font},sans-serif`

    return (
        <>
            {children}
        </>
    );
}