export interface MissionItem {
    id: string;
    label: string;
}
interface AiAssistantProps {
    title?: string;
    heading?: string;
    suggestions?: string[];
    placeholder?: string;
    onClose?: () => void;
    onMessage?: (text: string) => Promise<string>;
    allowUpload?: boolean;
    missions?: MissionItem[];
    selectedMissions?: string[];
    onMissionsChange?: (ids: string[]) => void;
}
export declare function AiAssistant({ title, heading, suggestions, placeholder, onClose, onMessage, allowUpload, missions, selectedMissions, onMissionsChange, }: AiAssistantProps): import("react/jsx-runtime").JSX.Element;
export {};
