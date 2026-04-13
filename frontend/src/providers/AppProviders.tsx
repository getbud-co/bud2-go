"use client";

import QueryProvider from "@/contexts/QueryContext";
import { LoggedUserProvider } from "@/contexts/LoggedUserContext";
import { OrganizationProvider } from "@/contexts/OrganizationContext";
import { ConfigDataProvider } from "@/contexts/ConfigDataContext";
import { ActivityDataProvider } from "@/contexts/ActivityDataContext";
import { PeopleDataProvider } from "@/contexts/PeopleDataContext";
import { MissionsDataProvider } from "@/contexts/MissionsDataContext";
import { SettingsDataProvider } from "@/contexts/SettingsDataContext";
import { IntegrationsDataProvider } from "@/contexts/IntegrationsDataContext";
import { SurveysDataProvider } from "@/contexts/SurveysDataContext";
import type { ReactNode } from "react";
import { usePathname } from "next/navigation";

interface AppProvidersProps {
  children: ReactNode;
  initialOrgId?: string;
}

export function AppProviders({ children, initialOrgId }: AppProvidersProps) {
  const pathname = usePathname();

  if (pathname === "/login") {
    return <>{children}</>;
  }

  return (
    <QueryProvider>
      <LoggedUserProvider>
        <OrganizationProvider initialOrgId={initialOrgId}>
          <ConfigDataProvider>
            <ActivityDataProvider>
              <PeopleDataProvider>
                <MissionsDataProvider>
                  <SurveysDataProvider>
                    <SettingsDataProvider>
                      <IntegrationsDataProvider>
                        {children}
                      </IntegrationsDataProvider>
                    </SettingsDataProvider>
                  </SurveysDataProvider>
                </MissionsDataProvider>
              </PeopleDataProvider>
            </ActivityDataProvider>
          </ConfigDataProvider>
        </OrganizationProvider>
      </LoggedUserProvider>
    </QueryProvider>
  );
}
