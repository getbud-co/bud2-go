interface ToastOptions {
    description?: string;
    duration?: number;
    action?: {
        label: string;
        onClick: () => void;
    };
}
export declare function toast(title: string, options?: ToastOptions): string;
export declare namespace toast {
    var success: (title: string, options?: ToastOptions) => string;
    var error: (title: string, options?: ToastOptions) => string;
    var warning: (title: string, options?: ToastOptions) => string;
    var black: (title: string, options?: ToastOptions) => string;
    var dismiss: (id?: string) => void;
}
export declare function Toaster(): import("react/jsx-runtime").JSX.Element | null;
export {};
