import React, { useEffect, useState } from 'react';
import { Table } from 'antd';
import axios from 'axios';

interface ContainerStatus {
  ip: string;
  pingTime: string;
  lastSuccess: string;
}

const ContainerTable: React.FC = () => {
  const [data, setData] = useState<ContainerStatus[]>([]);

  // Функция для получения данных с API
  const fetchData = async () => {
    try {
      // Запрашиваем данные с API pinger-service
      const response = await axios.get<ContainerStatus[]>('http://pinger-service:8081/pingIP/all');
      setData(response.data);
    } catch (error) {
      console.error('Error fetching container statuses:', error);
    }
  };

  // Используем useEffect для выполнения запроса при монтировании компонента
  useEffect(() => {
    fetchData(); // Первый запрос данных
    const interval = setInterval(fetchData, 5000); // Обновляем данные каждые 5 секунд

    // Очищаем интервал при размонтировании компонента
    return () => clearInterval(interval);
  }, []);

  // Определяем колонки для таблицы
  const columns = [
    {
      title: 'IP Address',
      dataIndex: 'ip',
      key: 'ip',
    },
    {
      title: 'Ping Time',
      dataIndex: 'pingTime',
      key: 'pingTime',
    },
    {
      title: 'Last Success',
      dataIndex: 'lastSuccess',
      key: 'lastSuccess',
    },
  ];

  // Возвращаем таблицу с данными
  return <Table dataSource={data} columns={columns} rowKey="ip" />;
};

export default ContainerTable;