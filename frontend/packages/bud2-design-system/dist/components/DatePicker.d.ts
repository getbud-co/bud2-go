import { ReactNode } from 'react';
import { CalendarDate } from './date-utils';
type MessageType = "error" | "attention" | "success";
type DatePickerSize = "sm" | "md" | "lg";
interface DatePickerBaseProps {
    /** Tamanho do trigger — combina com Button sizes (default: "md") */
    size?: DatePickerSize;
    label?: ReactNode;
    placeholder?: string;
    message?: string;
    messageType?: MessageType;
    disabled?: boolean;
    minDate?: CalendarDate;
    maxDate?: CalendarDate;
    className?: string;
}
interface SingleDatePickerProps extends DatePickerBaseProps {
    mode?: "single";
    value?: CalendarDate | null;
    defaultValue?: CalendarDate | null;
    onChange?: (date: CalendarDate | null) => void;
}
interface RangeDatePickerProps extends DatePickerBaseProps {
    mode: "range";
    value?: [CalendarDate | null, CalendarDate | null];
    defaultValue?: [CalendarDate | null, CalendarDate | null];
    onChange?: (range: [CalendarDate | null, CalendarDate | null]) => void;
}
export type DatePickerProps = SingleDatePickerProps | RangeDatePickerProps;
export declare function DatePicker(props: DatePickerProps): import("react/jsx-runtime").JSX.Element;
export {};
