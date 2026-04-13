type GoalStatus = "on-track" | "attention" | "off-track";
interface GoalProgressBarProps {
    label: string;
    value: number;
    target?: number;
    min?: number;
    formattedValue?: string;
    expected?: number;
    status?: GoalStatus;
    onChange?: (value: number) => void;
    className?: string;
}
export declare function GoalProgressBar({ label, value, target, min, formattedValue, expected, status: statusOverride, onChange, className, }: GoalProgressBarProps): import("react/jsx-runtime").JSX.Element;
type GaugeGoalType = "above" | "below" | "between";
interface GoalGaugeBarProps {
    label: string;
    value: number;
    low?: number;
    high?: number;
    goalType: GaugeGoalType;
    min?: number;
    max?: number;
    formattedValue?: string;
    status?: GoalStatus;
    onChange?: (value: number) => void;
    className?: string;
}
export declare function GoalGaugeBar({ label, value, low, high, goalType, min, max, formattedValue, status: statusOverride, onChange, className, }: GoalGaugeBarProps): import("react/jsx-runtime").JSX.Element;
export {};
