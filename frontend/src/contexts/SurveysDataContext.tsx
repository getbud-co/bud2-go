"use client";

import { createContext, useContext, useMemo, type ReactNode } from "react";
import {
  surveysSeed,
  type SurveyListItemData,
} from "@/lib/tempStorage/surveys-store";

interface SurveysDataContextValue {
  surveys: SurveyListItemData[];
}

const SurveysDataContext = createContext<SurveysDataContextValue | null>(null);

export function SurveysDataProvider({ children }: { children: ReactNode }) {
  const value = useMemo<SurveysDataContextValue>(
    () => ({ surveys: surveysSeed }),
    [],
  );

  return (
    <SurveysDataContext.Provider value={value}>
      {children}
    </SurveysDataContext.Provider>
  );
}

export function useSurveysData(): SurveysDataContextValue {
  const context = useContext(SurveysDataContext);

  if (!context) {
    throw new Error("useSurveysData must be used within a SurveysDataProvider");
  }

  return context;
}
