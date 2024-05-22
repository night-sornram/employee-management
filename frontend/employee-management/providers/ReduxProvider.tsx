"use client"
import React  from 'react'
import { store } from '@/store/store'
import { Provider } from 'react-redux'
import { persistStore } from 'redux-persist'
import { PersistGate } from 'redux-persist/integration/react'


type Props = { children: React.ReactNode }

export default function ReduxProvider({ children }: Props) {
    let persistor = persistStore(store)
  return (
    <Provider store={store}>
        <PersistGate loading={null} persistor={persistor}>
            {children}
        </PersistGate>
    </Provider>
  )
}