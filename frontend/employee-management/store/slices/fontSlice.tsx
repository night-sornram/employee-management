import { createSlice } from "@reduxjs/toolkit";
import { RootState } from "../store";
import { FontProp } from "@/interface";


const initialValue: FontProp = {
    font : "Inter"
};

const fontSlice = createSlice({
    name: "font",
    initialState: initialValue,
    reducers: { 
        updateFont: (state, action) => {
            state.font = action.payload.font
    },
    }
});

export default fontSlice.reducer;
export const { updateFont } = fontSlice.actions
export const appSelector = (state:RootState) => state.appReducer;