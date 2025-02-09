
import React, { useEffect, useState } from 'react';
import { Table } from 'antd'; // Компонент Table из библиотеки antd для отображения таблицы
import axios from 'axios'; // Библиотека для выполнения HTTP-запросов

// Интерфейс для описания структуры данных о состоянии контейнера
interface ContainerStatus {
    ip: string; // IP-адрес контейнера
    pingTime: string; // Время последнего пинга
    lastSuccess: string; // Дата последней успешной попытки пинга
}

// Основной компонент ContainerTable, который отображает таблицу с данными о контейнерах
const ContainerTable: React.FC = () => {
    // Используем хук useState для хранения данных о состоянии контейнеров
    const [data, setData] = useState<ContainerStatus[]>([]);

    // Используем хук useEffect для выполнения side-эффектов (в данном случае — запрос данных)
    useEffect(() => {
        // Асинхронная функция для получения данных о состоянии контейнеров
        const fetchData = async () => {
            try {
                // Выполняем GET-запрос к API backend для получения данных
                const response = await axios.get<ContainerStatus[]>('http://localhost:8080/ping');
                // Обновляем состояние data с полученными данными
                setData(response.data);
            } catch (error) {
                // В случае ошибки выводим её в консоль
                console.error('Error fetching container statuses:', error);
            }
        };

        // Вызываем fetchData сразу при монтировании компонента
        fetchData();
        // Устанавливаем интервал для периодического обновления данных
        const interval = setInterval(fetchData, 5000);

        return () => clearInterval(interval);
    }, []);

    // Определяем колонки для таблицы
    const columns = [
        {
            title: 'IP Address', // Заголовок колонки
            dataIndex: 'ip', // Поле данных, которое будет отображаться в этой колонке
            key: 'ip', // Уникальный ключ для колонки
        },
        {
            title: 'Ping Time', // Заголовок колонки
            dataIndex: 'pingTime', // Поле данных для времени пинга
            key: 'pingTime', // Уникальный ключ для колонки
        },
        {
            title: 'Last Success', // Заголовок колонки
            dataIndex: 'lastSuccess', // Поле данных для даты последней успешной попытки
            key: 'lastSuccess', // Уникальный ключ для колонки
        },
    ];

    // Возвращаем компонент Table с данными и колонками
    return <Table dataSource={data} columns={columns} rowKey="ip" />;
};

export default ContainerTable;