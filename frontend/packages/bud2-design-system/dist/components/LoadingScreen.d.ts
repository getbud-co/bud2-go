export interface LoadingScreenProps {
    /** Mensagem exibida abaixo do logo (default: "Carregando...") */
    message?: string;
    /** Classe CSS adicional */
    className?: string;
}
/**
 * Tela de loading fullscreen com animação liquid fill no logo do Bud.
 *
 * Usada no boot da aplicação enquanto os dados iniciais são carregados.
 * A animação preenche o símbolo do Bud com um sweep left→right,
 * simulando um líquido com borda ondulada via cubic beziers.
 *
 * @example
 * ```tsx
 * import { LoadingScreen } from "@getbud-co/buds";
 *
 * // Uso padrão
 * <LoadingScreen />
 *
 * // Com mensagem customizada
 * <LoadingScreen message="Conectando ao servidor..." />
 * ```
 */
export declare function LoadingScreen({ message, className, }: LoadingScreenProps): import("react/jsx-runtime").JSX.Element;
