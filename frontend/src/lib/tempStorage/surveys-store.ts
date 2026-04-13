export interface SurveyListItemData {
  id: string;
  name: string;
  type: string;
  status: string;
  endDate: string;
  totalResponses: number;
  totalRecipients: number;
  completionRate: number;
  managerIds: string[];
}

export const surveysSeed: SurveyListItemData[] = [];
