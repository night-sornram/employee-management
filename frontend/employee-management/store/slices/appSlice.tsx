import { createSlice } from "@reduxjs/toolkit";
import { RootState } from "../store";
import { StateProp } from "@/interface";


const initialValue: StateProp = {
    notification: "all",
    email: false,
};

const appSlice = createSlice({
    name: "app",
    initialState: initialValue,
    reducers: { 
        updateNotification: (state, action) => {
            state.notification = action.payload.notification;
            state.email = action.payload.email;
        }
    
    },
});

export default appSlice.reducer;
export const { updateNotification } = appSlice.actions;
export const appSelector = (state:RootState) => state.appReducer;