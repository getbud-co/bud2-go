import { ReactNode } from 'react';
export interface TabItem {
    /** Identificador único da aba */
    value: string;
    /** Texto exibido na aba */
    label: string;
    /** Conteúdo extra ao lado do label (ex: Badge com contagem) */
    badge?: ReactNode;
    /** Desabilita a aba */
    disabled?: boolean;
}
export interface TabBarProps {
    /** Lista de abas */
    tabs: TabItem[];
    /** Valor da aba ativa */
    activeTab: string;
    /** Callback ao trocar de aba */
    onTabChange: (value: string) => void;
    /** Label acessível para o tablist */
    ariaLabel?: string;
    /** ID base para construir IDs de tab/panel (ex: "pesquisa" → tab "pesquisa-tab-resumo", panel "pesquisa-panel-resumo") */
    id?: string;
    className?: string;
}
/** Gera o ID do tab button a partir do baseId e value */
export declare function getTabId(baseId: string, value: string): string;
/** Gera o ID do tabpanel a partir do baseId e value */
export declare function getPanelId(baseId: string, value: string): string;
export declare function TabBar({ tabs, activeTab, onTabChange, ariaLabel, id, className, }: TabBarProps): import("react/jsx-runtime").JSX.Element;
