export interface ScaleInputProps {
    /** Minimum value (default 0) */
    min?: number;
    /** Maximum value (default 10) */
    max?: number;
    /** Currently selected value */
    value?: number;
    /** Called when a value is selected */
    onChange?: (value: number) => void;
    /** Label displayed below min value */
    minLabel?: string;
    /** Label displayed below max value */
    maxLabel?: string;
    /** Disables all buttons */
    disabled?: boolean;
    /** Visual size: sm = compact, md = default */
    size?: "sm" | "md";
    className?: string;
}
export declare function ScaleInput({ min, max, value, onChange, minLabel, maxLabel, disabled, size, className, }: ScaleInputProps): import("react/jsx-runtime").JSX.Element;
