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
    /** Called when user clicks "Usar como base" on an assistant message */
    onUseAsBase?: (messageContent: string) => void;
    allowUpload?: boolean;
    missions?: MissionItem[];
    selectedMissions?: string[];
    onMissionsChange?: (ids: string[]) => void;
}
export declare function AiAssistant({ title, heading, suggestions, placeholder, onClose, onMessage, onUseAsBase, allowUpload, missions, selectedMissions, onMissionsChange, }: AiAssistantProps): import("react/jsx-runtime").JSX.Element;
export {};
