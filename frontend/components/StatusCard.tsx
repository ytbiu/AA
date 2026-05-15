import React from 'react';

type StatusCardType = 'text' | 'address' | 'balance' | 'boolean';

interface StatusCardProps {
  title: string;
  value: string | boolean | number;
  type: StatusCardType;
}

const StatusCard: React.FC<StatusCardProps> = ({ title, value, type }) => {
  const formatValue = () => {
    switch (type) {
      case 'address':
        // Format address with ellipsis in the middle
        const stringValue = String(value);
        if (stringValue.length > 10) {
          return `${stringValue.substring(0, 6)}...${stringValue.substring(stringValue.length - 4)}`;
        }
        return stringValue;
      case 'balance':
        // Format balance with 4 decimal places
        return typeof value === 'number' ? value.toFixed(4) : String(value);
      case 'boolean':
        // Show boolean as Yes/No
        return value ? 'Yes' : 'No';
      case 'text':
      default:
        return String(value);
    }
  };

  return (
    <div className="bg-white rounded-lg shadow p-4">
      <h3 className="text-sm font-medium text-gray-500">{title}</h3>
      <p className="mt-1 text-lg font-semibold text-gray-900">{formatValue()}</p>
    </div>
  );
};

export default StatusCard;