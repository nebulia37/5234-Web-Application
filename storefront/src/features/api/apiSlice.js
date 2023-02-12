import { createSlice } from '@reduxjs/toolkit'

export const apiSlice = createSlice({
    name: "api",
    initialState: {
        inventory: process.env.REACT_APP_INVENTORY_API_SERVER? process.env.REACT_APP_INVENTORY_API_SERVER: "http://localhost:9000/",
        order: process.env.REACT_APP_ORDER_API_SERVER? process.env.REACT_APP_ORDER_API_SERVER: "http://localhost:9001/",
        payment: process.env.REACT_APP_PAYMENT_API_SERVER? process.env.REACT_APP_PAYMENT_API_SERVER: "http://localhost:9003/",
    },
    reducers: {
        update: (state, action) => {
            state.inventory = action.payload
        },
    }
});

export const { update } = apiSlice.actions;

export default apiSlice.reducer;